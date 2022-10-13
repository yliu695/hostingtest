import { Typography, Input, Form, Button, InputNumber, Message } from '@arco-design/web-react';
import { notifyContact } from 'utils/request';
const { Title, Paragraph, Text } = Typography;
const TextArea = Input.TextArea;
const FormItem = Form.Item;

export function Contact () {
  const [form] = Form.useForm();

  return (
    <Typography style={{ marginTop: 10 }}>
      <Title heading={4}>Contact With US</Title>
      <Form form={form} style={{ width: 600 }} autoComplete='off'>
        <FormItem label='Your Name' field='name' rules={[{ required: true }]}>
          <Input required type='text' placeholder='please enter your name...' />
        </FormItem>
        <FormItem label='Email' field='email' rules={[{ required: true, type: 'email' }]}>
          <Input required type='email' placeholder='please enter your email...' />
        </FormItem>
        <FormItem label='Feedback' field='feedback' rules={[{ required: true, minLength: 10 }]}>
          <TextArea required minLength={10} placeholder='Please enter ...' style={{ minHeight: 64, width: 350 }} />
        </FormItem>
        <FormItem wrapperCol={{ offset: 5 }}>
          <Button type='primary' onClick={() => {
            form.validate().then(values => {
              const { name, email, feedback } = values;
              notifyContact(name, email, feedback).then(res => {
                if (res.code == 0) {
                  Message.success('Submit success');
                }
              })
            })
            form.getFields()
          }}>Submit</Button>
        </FormItem>
      </Form>
    </Typography>
  );
}