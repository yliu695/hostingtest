import { Grid, Link, Typography, Space } from '@arco-design/web-react';
const Row = Grid.Row;
const Col = Grid.Col;
const { Title, Paragraph, Text } = Typography;


export function Footer () {
    return (
        <div style={{
            width: '100%',
            backgroundColor: 'rgb(var(--arcoblue-1))',
            marginTop: 15,
        }}>
            <Row style={{
                marginLeft: 100,
                marginRight: 15,
            }}>
                <Col flex={3}>
                    <Typography style={{ marginTop: 10 }}>
                        <Title heading={5}>About US</Title>
                        <Paragraph>
                            Team 26
                        </Paragraph>
                    </Typography>
                </Col>
                <Col flex={1}>
                </Col>
                <Col flex={3}>
                    <Typography style={{ marginTop: 10 }}>
                        <Title heading={5}>Resources</Title>
                        <Space direction='horizontal'>
                            <Link href='/news'> News </Link>
                            <Link href='/events'> Events </Link>
                            <Link href='/staffs'> Staffs </Link>
                            <Link href='/phds'> PhD Students </Link>
                            <Link href='/projects'> Projects </Link>
                            <Link href='/resuorces'> Resources </Link>
                            <Link href='/cpntact'> Contact Us </Link>
                        </Space>
                    </Typography>
                </Col>
            </Row>
        </div>
    );
};