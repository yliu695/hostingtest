import React, {
  useState,
  useEffect,
} from 'react';
import { List, Avatar, Grid, Link, Button, Image, Message } from '@arco-design/web-react';
import { checkIsAdminLogin, phdList } from 'utils/request';
const Row = Grid.Row;
const Col = Grid.Col;

export function PhDs () {
  const [phds, setPhds] = useState([]);


  const [isAdminLogin, setIsAdminLogin] = useState(false);
  useEffect(() => {
    checkIsAdminLogin().then(res => {
      if (res.code == 0 && res.data.isLogin) {
        setIsAdminLogin(true);
      }
    })
  }, [])

  useEffect(() => {
    phdList().then(res => {
      if (res.code != 0) {
        Message.error(res.msg);
        return
      }

      setPhds(res.data.data)
    })
  }, []);

  return (
    <>
      {isAdminLogin && <Row style={{ marginBottom: 16 }}>
        <Col flex={10}>
        </Col>
        <Col flex={2}>
          <Link href='/phdsManage'>
            <Button style={{
              marginTop: 10,
            }} type='primary' size='small'>Manage Phds</Button>
          </Link>
        </Col>
      </Row>
      }
      <List
        wrapperStyle={{ marginLeft: 20, marginRight: 200 }}
        bordered={false}
        dataSource={phds}
        render={(item, index) => (
          <List.Item
            key={index}
            style={{ padding: '20px 0', borderBottom: '1px solid var(--color-fill-3)' }}
          >
            <List.Item.Meta
              avatar={
                <Image
                  width={200}
                  src={item.avatar}
                  footerPosition='outer'
                  alt='picture'
                />
              }
              title={<>
                {item.name} - {item.job}
              </>}
              description={
                <div dangerouslySetInnerHTML={{ __html: item.intro }} />
              }
            />
          </List.Item>
        )}
      />
    </>
  );
}