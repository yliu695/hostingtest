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
import { checkIsAdminLogin, eventList } from 'utils/request';
import { dateFormat } from 'utils/util';
const { Title, Paragraph } = Typography;
const MenuItem = Menu.Item;
const Row = Grid.Row;
const Col = Grid.Col;
const { Meta } = Card;
export function Events () {
  const [events, setEvents] = useState([]);
  const [event, setEvent] = useState({})
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
    eventList().then(res => {
      if (res.code != 0) {
        Message.error(res.msg);
        return
      }

      res.data.data.forEach((item, index) => {
        // parse tags
        item.tags = item.tags.split('|').filter(item => item != '');
        res.data.data[index] = item;

        if (item.id == params.get('id')) {
          // get event by special id
          setEvent(item);
        } else if (params.get('id') == undefined && (index == 0 || new Date(item.event_time * 1000) > new Date())) {
          // get first event
          setEvent(item);
        }
      });

      setEvents(res.data.data)
    })
  }, []);

  return (
    <>
      {isAdminLogin &&
        <Row style={{ marginBottom: 16 }}>
          <Col flex={10}>
          </Col>
          <Col flex={2}>
            <Link href='/eventsManage'>
              <Button style={{
                marginTop: 10,
              }} type='primary' size='small'>Manage Events</Button>
            </Link>
          </Col>
        </Row>
      }

      <Typography style={{
        marginLeft: '100px',
        marginRight: '100px',
      }}>
        <Title heading={1}>{event.title}</Title>
        <Title heading={6}>{dateFormat(new Date(event.event_time * 1000))} </Title>
        <br />
        <Space size='medium' style={{
          marginTop: '15px'
        }}>
          {event.tags && event.tags.map((tag, idx) => (
            <Tag color='arcoblue' key={idx} defaultChecked>
              {tag}
            </Tag>
          ))}
        </Space>

        <br />
        <div dangerouslySetInnerHTML={{ __html: event.content }} />

        <Title heading={2}>More Events</Title>
        <Row style={{
          marginBottom: 16,
          justify: 'start'
        }}>
          {events.map((ev, _) => {
            if (ev.id != event.id) {
              return (
                <Col flex={'auto'}>
                  <Link href={`/events?id=${ev.id}`}>
                    <Card
                      onClick={() => {
                        Link.go(`/events?id=${ev.id}`)
                      }}
                      hoverable
                      style={{ width: 300 }}
                      cover={
                        <div>
                          <img
                            style={{
                              width: '100%',
                            }}
                            src={ev.cover || "/events.jpeg"}
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