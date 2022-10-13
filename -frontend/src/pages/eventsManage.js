import React, {
    useState,
    useEffect,
} from 'react';
import { Button, Message, Table, Input, Select, Form, Space, Tag, Trigger, Typography, DatePicker } from '@arco-design/web-react';
import { addEvent, checkIsAdminLogin, deleteEvent, editEvent, eventList } from 'utils/request';
import { WYSISYGEditor } from 'components/editor';
import { EditableCell } from 'components/editableRowCell';
import { EditableRow } from 'components/editableRowCell';
export function EventsManage () {
    const [count, setCount] = useState(5);
    const [data, setData] = useState([]);

    const columns = [
        {
            title: 'Title',
            dataIndex: 'title',
            editable: true
        },
        {
            title: 'Content',
            dataIndex: 'content',
            editable: false,
            render: (_, record) => (
                <Trigger
                    popup={() => (
                        <div style={{
                            width: '90%',
                            backgroundColor: 'white',
                            maxHeight: '80%',
                            border: '1px solid #ccc',
                            margin: '0 auto',
                        }}>
                            <WYSISYGEditor html={record.content} onChange={html => {
                                record.content = html
                            }} />
                        </div>
                    )}
                    onVisibleChange={visible => {
                        // save when popup close
                        if (!visible) {
                            handleSave(record)
                        }
                    }}

                    mouseEnterDelay={400}
                    mouseLeaveDelay={400}
                    trigger='click'
                >
                    <Button type='primary' size='mini'>
                        Edit
                    </Button>
                </Trigger>
            ),
        },
        {
            title: 'Cover',
            dataIndex: 'cover',
            editable: true
        },
        {
            title: 'Tags',
            dataIndex: 'tags',
            type: 'tags',
            editable: true,
            render: (_, record) => (
                <Space size='medium'>
                    {
                        record.tags && record.tags.map((tag, idx) => (
                            <Tag color='arcoblue' key={idx} defaultChecked>
                                {tag}
                            </Tag>
                        ))
                    }
                </Space>
            ),
        },
        {
            title: 'Event Time',
            dataIndex: 'event_time',
            editable: false,
            render: (_, record) => (
                <>
                    <DatePicker
                        showTime
                        allowClear={false}
                        // timezone={Intl.DateTimeFormat().resolvedOptions().timeZone}
                        value={new Date(record.event_time * 1000)}
                        onOk={(v, vd) => {
                            if (vd) {
                                record.event_time = vd.toDate().getTime() / 1000
                            }
                            handleSave(record)
                        }}
                    />
                </>
            ),
        },
        {
            title: 'Operation',
            dataIndex: 'op',
            render: (_, record) => (
                <>
                    <Space>
                        <Button
                            onClick={() => removeRow(record)}
                            type='primary'
                            size='mini'
                            status='danger'
                        >
                            Delete
                        </Button>
                    </Space>
                </>
            )
        }
    ];

    let parseEvent = (event) => {
        event.key = `old_${event.id}`;
        event.tags = event.tags ? event.tags.split('|').filter(item => item != '') : [];
        event.cover = event.cover || '/events.jpeg';
        return event
    }

    useEffect(() => {
        eventList().then(res => {
            if (res.code != 0) {
                Message.error(res.msg);
                return
            }

            // set key for each row
            res.data.data.forEach((item, index) => {
                res.data.data[index] = parseEvent(item);
            });
            setData(res.data.data);
        })
    }, []);

    function handleSave (row) {
        let newRow = row

        let update = () => {
            newRow.new = false

            const newData = [...data]
            const index = newData.findIndex(item => row.key === item.key);
            newData.splice(index, 1, { ...newData[index], ...newRow });
            setData(newData);
        }

        if (row.new) {
            // create new
            addEvent(row.title, row.content, row.tags, row.cover, row.event_time).then(res => {
                if (res.code != 0) {
                    Message.error(res.msg);
                    return
                }
                newRow = parseEvent(res.data)
                update()
            })
        } else {
            // update event
            editEvent(row.id, row.title, row.content, row.tags, row.cover, row.event_time).then(res => {
                if (res.code != 0) {
                    Message.error(res.msg);
                    return
                }
            })
            update()
        }
    }

    function removeRow (row) {
        deleteEvent(row.id).then(res => {
            if (res.code != 0) {
                Message.error(res.msg);
                return
            }
        })

        setData(data.filter(item => item.key !== row.key));
    }

    function addRow () {
        setCount(count + 1);
        setData(
            data.concat({
                new: true,
                key: `new_${count + 1}`,
                title: 'new event',
                content: '',
                tags: [],
                event_time: new Date().getTime() / 1000,
                cover: '/events.jpeg',
            })
        );
    }

    return (
        <div style={{
            marginLeft: 40,
            marginRight: 40,
        }}>
            <Button
                style={{
                    marginBottom: 10
                }}
                type='primary'
                onClick={addRow}
            >
                Add
            </Button>
            <Table
                data={data}
                components={{
                    body: {
                        row: EditableRow,
                        cell: EditableCell
                    }
                }}
                columns={columns.map(column =>
                    column.editable
                        ? {
                            ...column,
                            onCell: () => ({
                                onHandleSave: handleSave
                            })
                        }
                        : column
                )}
            />
        </div>
    );
}
