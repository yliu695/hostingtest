import '@wangeditor/editor/dist/css/style.css' // import css

import React, { useState, useEffect } from 'react'
import { Editor, Toolbar } from '@wangeditor/editor-for-react'
import { IDomEditor, IEditorConfig, IToolbarConfig } from '@wangeditor/editor'

export function WYSISYGEditor (props) {
    // editor instance
    const [editor, setEditor] = useState(null)

    // editor content
    const [html, setHtml] = useState(props.html)

    const toolbarConfig = {
        excludeKeys: ['uploadVideo']
    }

    const editorConfig = {
        placeholder: 'Type here...',
        MENU_CONF: {
            uploadImage: {
                // server: '/api/upload',
                base64LimitSize: 5 * 1024 * 1024 * 1024, // 5G
            }
        }
    }

    // Timely destroy editor, important!
    useEffect(() => {
        return () => {
            if (editor == null) return
            editor.destroy()
            setEditor(null)
        }
    }, [editor])

    return (
        <>
            <div style={{ border: '1px solid #ccc', zIndex: 100 }}>
                <Toolbar
                    editor={editor}
                    defaultConfig={toolbarConfig}
                    mode="default"
                    style={{ borderBottom: '1px solid #ccc' }}
                />
                <Editor
                    defaultConfig={editorConfig}
                    value={html}
                    onCreated={setEditor}
                    onChange={editor => {
                        // console.log(editor.getConfig().toolbarKeys)
                        setHtml(editor.getHtml())
                        if (props.onChange) {
                            props.onChange(editor.getHtml())
                        }
                    }}
                    mode="default"
                    style={{ height: '500px', overflowY: 'hidden' }}
                />
            </div>
        </>
    )
}