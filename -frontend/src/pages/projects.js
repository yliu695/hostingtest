import React, {
  useState,
  useEffect,
} from 'react';
import { List, Avatar, Grid, Link, Button, Image, Message } from '@arco-design/web-react';
import { checkIsAdminLogin, projectList } from 'utils/request';
const Row = Grid.Row;
const Col = Grid.Col;

export function Projects () {
  const [projects, setProjects] = useState([]);


  const [isAdminLogin, setIsAdminLogin] = useState(false);
  useEffect(() => {
    checkIsAdminLogin().then(res => {
      if (res.code == 0 && res.data.isLogin) {
        setIsAdminLogin(true);
      }
    })
  }, [])

  useEffect(() => {
    projectList().then(res => {
      if (res.code != 0) {
        Message.error(res.msg);
        return
      }

      setProjects(res.data.data)
    })
  }, []);

  return (
    <>
      {isAdminLogin && <Row style={{ marginBottom: 16 }}>
        <Col flex={10}>
        </Col>
        <Col flex={2}>
          <Link href='/projectsManage'>
            <Button style={{
              marginTop: 10,
            }} type='primary' size='small'>Manage Projects</Button>
          </Link>
        </Col>
      </Row>
      }
      <List
        wrapperStyle={{ marginLeft: 20, marginRight: 200 }}
        bordered={false}
        dataSource={projects}
        render={(item, index) => (
          <List.Item
            key={index}
            style={{ padding: '20px 0', borderBottom: '1px solid var(--color-fill-3)' }}
          >
            <List.Item.Meta
              title={<>
                {item.name}
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