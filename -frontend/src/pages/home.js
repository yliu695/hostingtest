import React, {
  useState,
  useEffect,
} from 'react';
import { Carousel, Typography, Image, Space, Link, Message } from '@arco-design/web-react';
import { eventList, newsList } from 'utils/request';
const newsMap = [
  {
    name: 'News 1',
    img: '//p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/cd7a1aaea8e1c5e3d26fe2591e561798.png~tplv-uwbnlip3yd-webp.webp',
    desc: 'News 1 description',
    link: '/news',
  },
  {
    name: 'Events 1',
    img: '//p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/6480dbc69be1b5de95010289787d64f1.png~tplv-uwbnlip3yd-webp.webp',
    desc: 'Events 1 description',
    link: '/events',
  },
  {
    name: 'News 2',
    img: '//p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/0265a04fddbd77a19602a15d9d55d797.png~tplv-uwbnlip3yd-webp.webp',
    desc: 'News 2 description',
    link: '/news',
  },
  {
    name: 'Events 2',
    img: '//p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/24e0dd27418d2291b65db1b21aa62254.png~tplv-uwbnlip3yd-webp.webp',
    desc: 'Events 4 description',
    link: '/events',
  }
];

export function Home () {
  const srcList = [
    '//p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/a8c8cdb109cb051163646151a4a5083b.png~tplv-uwbnlip3yd-webp.webp',
    '//p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/e278888093bef8910e829486fb45dd69.png~tplv-uwbnlip3yd-webp.webp',
    '//p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/3ee5f13fb09879ecb5185e440cef6eb9.png~tplv-uwbnlip3yd-webp.webp',
    '//p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/8361eeb82904210b4f55fab888fe8416.png~tplv-uwbnlip3yd-webp.webp',
  ];

  const [events, setEvents] = useState([])
  const [news, setNews] = useState([])

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
      });

      setEvents(res.data.data)
    })

    newsList().then(res => {
      if (res.code != 0) {
        Message.error(res.msg);
        return
      }

      res.data.data.forEach((item, index) => {
        // parse tags
        item.tags = item.tags.split('|').filter(item => item != '');
        res.data.data[index] = item;
      });

      setNews(res.data.data)
    })
  }, []);

  return (
    <>
      <Carousel
        autoPlay
        animation='card'
        showArrow='never'
        indicatorPosition='outer'
        style={{ width: '100%', marginTop: '10px' }}
      >
        {news.map((n, index) => {
          if (index > 1) {
            return null;
          }

          return (
            <div
              key={index}
              style={{ width: '60%' }}
            >
              <Link href={`/news?id=${n.id}`}>
                <Image style={{
                  width: '80%',
                  height: 350
                }} preview={false} src={n.cover || "/news.jpeg"} title={n.title} />
              </Link>
            </div>
          )
        })}

        {events.map((event, index) => {
          if (index > 1) {
            return null;
          }

          return (
            <div
              key={index}
              style={{ width: '60%' }}
            >
              <Link href={`/events?id=${event.id}`}>
                <Image style={{
                  width: '80%',
                  height: 350
                }} preview={false} src={event.cover || "/events.jpeg"} title={event.title} />
              </Link>
            </div>
          )
        })}
      </Carousel>
    </>
  );
}