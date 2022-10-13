import React, {
    useState,
    useRef,
    useEffect,
    useContext,
    useCallback
} from 'react';
import { Button, Message, DatePicker, Table, Input, Select, Form, Space, Tag, Trigger, Typography } from '@arco-design/web-react';
import { IconPlus } from '@arco-design/web-react/icon';

const FormItem = Form.Item;
const EditableContext = React.createContext({});

export function EditableCell (props) {
    const { children, className, rowData, column, onHandleSave } = props;
    const ref = useRef(null);
    const refInput = useRef(null);
    const { getForm } = useContext(EditableContext);
    const [editing, setEditing] = useState(false);
    const handleClick = useCallback(
        e => {
            if (
                editing &&
                column.editable &&
                ref.current &&
                !ref.current.contains(e.target)
            ) {
                cellValueChangeHandler(rowData[column.dataIndex]);
            }
        },
        [editing, rowData, column]
    );
    useEffect(() => {
        editing && refInput.current && refInput.current.focus();
    }, [editing]);
    useEffect(() => {
        document.addEventListener('click', handleClick, true);
        return () => {
            document.removeEventListener('click', handleClick, true);
        };
    }, [handleClick]);

    const cellValueChangeHandler = value => {
        const form = getForm();
        form.validate([column.dataIndex], (errors, values) => {
            if (!errors || !errors[column.dataIndex]) {
                setEditing(!editing);
                onHandleSave && onHandleSave({ ...rowData, ...values });
            }
        });
    };

    if (editing) {
        // tags
        if (column.type === 'tags') {
            return (
                <div ref={ref} >
                    <TagsCellEdit rowData={rowData} column={column} />
                </div>
            )
        }

        return (
            <div ref={ref}>
                <FormItem
                    style={{
                        marginBottom: 0
                    }}
                    labelCol={{
                        span: 0
                    }}
                    wrapperCol={{
                        span: 24
                    }}
                    initialValue={rowData[column.dataIndex]}
                    field={column.dataIndex}
                    rules={[
                        {
                            required: true
                        }
                    ]}
                >
                    <Input ref={refInput} onPressEnter={cellValueChangeHandler} />
                </FormItem>
            </div>
        );
    }

    return (
        <div
            className={column.editable ? `editable-cell ${className}` : className}
            onClick={() => column.editable && setEditing(!editing)}
        >
            {children}
        </div>
    );
}

function TagsCellEdit (props) {
    const { rowData, column } = props;
    const [tags, setTags] = useState(rowData[column.dataIndex]);
    const [showInput, setShowInput] = useState(false);
    const [inputValue, setInputValue] = useState('');

    function addTag () {
        if (inputValue) {
            tags.push(inputValue);
            setTags(tags);
            setInputValue('');
        }

        setShowInput(false);
    }

    function removeTag (removeTag) {
        const newTags = tags.filter((tag) => tag !== removeTag);
        setTags(newTags);
    }

    return (
        <Space size={20}>
            {tags.map((tag, index) => {
                return (
                    <Tag
                        key={tag}
                        closable={index !== 0}
                        onClose={() => removeTag(tag)}
                    >
                        {tag}
                    </Tag>
                );
            })}
            {showInput ? (
                <Input
                    autoFocus
                    size='mini'
                    value={inputValue}
                    style={{ width: 84 }}
                    onPressEnter={addTag}
                    onBlur={addTag}
                    onChange={setInputValue}
                />
            ) : (
                <Tag
                    icon={<IconPlus />}
                    style={{
                        width: 84,
                        backgroundColor: 'var(--color-fill-2)',
                        border: '1px dashed var(--color-fill-3)',
                        cursor: 'pointer',
                    }}
                    onClick={() => setShowInput(true)}
                >
                    Add Tag
                </Tag>
            )}
        </Space>
    )
}

export function EditableRow (props) {
    const { children, record, className, ...rest } = props;
    const refForm = useRef(null);

    const getForm = () => refForm.current;

    return (
        <EditableContext.Provider
            value={{
                getForm
            }}
        >
            <Form
                style={{
                    display: 'table-row'
                }}
                children={children}
                ref={refForm}
                wrapper='tr'
                wrapperProps={rest}
                className={`${className} editable-row`}
            />
        </EditableContext.Provider>
    );
}