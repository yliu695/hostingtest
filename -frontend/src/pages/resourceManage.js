import React, {
    useState,
    useEffect,
} from 'react';
import { Button, Message, Table, Input, Select, Form, Space, Tag, Trigger, Typography } from '@arco-design/web-react';
import { addNews, addResource, deleteNews, deleteResource, editNews, editResource, newsList, resourceList } from 'utils/request';
import { WYSISYGEditor } from 'components/editor';
import { EditableCell } from 'components/editableRowCell';
import { EditableRow } from 'components/editableRowCell';
export function ResourcesManage () {
    const [count, setCount] = useState(5);
    const [data, setData] = useState([]);

    const columns = [
        {
            title: 'Name',
            dataIndex: 'name',
            editable: true
        },
        {
            title: 'Intro',
            dataIndex: 'intro',
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
                            <WYSISYGEditor html={record.intro} onChange={html => {
                                record.intro = html
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
            title: 'Link',
            dataIndex: 'link',
            editable: true
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

    let parseResource = (resource) => {
        resource.key = `old_${resource.id}`;
        return resource
    }

    useEffect(() => {
        resourceList().then(res => {
            if (res.code != 0) {
                Message.error(res.msg);
                return
            }

            // set key for each row
            res.data.data.forEach((item, index) => {
                res.data.data[index] = parseResource(item);
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
            addResource(row.name, row.intro, row.link).then(res => {
                if (res.code != 0) {
                    Message.error(res.msg);
                    return
                }
                newRow = parseResource(res.data)
                update()
            })
        } else {
            // update resource
            editResource(row.id, row.name, row.intro, row.link).then(res => {
                if (res.code != 0) {
                    Message.error(res.msg);
                    return
                }
            })
            update()
        }
    }

    function removeRow (row) {
        deleteResource(row.id).then(res => {
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
                name: 'new resource',
                intro: 'intro',
                link: 'https://www.example.com'
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
