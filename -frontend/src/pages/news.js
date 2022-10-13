import React, {
  useState,
  useEffect,
} from 'react';

import {
  useSearchParams
} from 'react-router-dom';
import { Typography, Space, Image, Card, Tag, Grid, Link, Button, Message } from '@arco-design/web-react';
import { Menu, Trigger } from '@arco-design/web-react';
import { IconMessage, IconClose, IconBug, IconBulb } from '@arco-design/web-react/icon';
import { checkIsAdminLogin, newsList } from 'utils/request';
import { dateFormat } from 'utils/util';
const { Title, Paragraph } = Typography;
const MenuItem = Menu.Item;
const Row = Grid.Row;
const Col = Grid.Col;
const { Meta } = Card;
export function News () {
  const [news, setNews] = useState([]);
  const [theNews, setTheNews] = useState({})
  const [params, setParams] = useSearchParams();


  const [isAdminLogin, setIsAdminLogin] = useState(false);
  useEffect(() => {
    checkIsAdminLogin().then(res => {
      if (res.code == 0 && res.data.isLogin) {
        setIsAdminLogin(true);
      }
    })
  }, [])

  useEffect(() => {
    newsList().then(res => {
      if (res.code != 0) {
        Message.error(res.msg);
        return
      }

      res.data.data.forEach((item, index) => {
        // parse tags
        item.tags = item.tags.split('|').filter(item => item != '');
        res.data.data[index] = item;

        if (item.id == params.get('id')) {
          // get news by special id
          setTheNews(item);
        } else if (params.get('id') == undefined && index == 0) {
          // get first news
          setTheNews(item);
        }
      });

      setNews(res.data.data)
    })
  }, []);

  return (
    <>
      {isAdminLogin &&
        <Row style={{ marginBottom: 16 }}>
          <Col flex={10}>
          </Col>
          <Col flex={2}>
            <Link href='/newsManage'>
              <Button style={{
                marginTop: 10,
              }} type='primary' size='small'>Manage News</Button>
            </Link>
          </Col>
        </Row>
      }

      <Typography style={{
        marginLeft: '100px',
        marginRight: '100px',
      }}>
        <Title heading={1}>{theNews.title}</Title>
        <Title heading={6}>{dateFormat(new Date(theNews.create_time))}</Title>
        <br />
        <Space size='medium' style={{
          marginTop: '15px'
        }}>
          {theNews.tags && theNews.tags.map((tag, idx) => (
            <Tag key={idx} color='arcoblue' defaultChecked>
              {tag}
            </Tag>
          ))}
        </Space>

        <br />
        <div dangerouslySetInnerHTML={{ __html: theNews.content }} />

        <Title heading={2}>More News</Title>
        <Row style={{
          marginBottom: 16,
          justify: 'start'
        }}>
          {news.map((ev, _) => {
            if (ev.id != theNews.id) {
              return (
                <Col flex={'auto'} key={ev.id}>
                  <Link href={`/news?id=${ev.id}`}>
                    <Card
                      onClick={() => {
                        Link.go(`/news?id=${ev.id}`)
                      }}
                      hoverable
                      style={{ width: 300 }}
                      cover={
                        <div>
                          <img
                            style={{
                              width: '100%',
                            }}
                            src={ev.cover || "/news.jpeg"}
                          />
                        </div>
                      }
                    >
                      <Meta style={{
                        overflow: 'w'
                      }}
                        title={ev.title}
                      />
                    </Card>
                  </Link>
                </Col>
              )
            }
          })
          }
        </Row>
      </Typography>
    </>
  );
}