import React, {
  useState,
  useEffect,
} from 'react';
import './App.css';
import { Avatar, Menu, Button, Trigger, Message } from '@arco-design/web-react';
import "@arco-design/web-react/dist/css/arco.css";
import '@wangeditor/editor/dist/css/style.css';
import { Link, Route, Routes, BrowserRouter } from "react-router-dom";
import { Home } from 'pages/home'
import { Contact } from 'pages/contact'
import { Events } from 'pages/events'
import { News } from 'pages/news'
import { Projects } from 'pages/projects'
import { Resources } from 'pages/resuorces'
import { Staffs } from 'pages/staffs'
import { Footer } from 'components/footer';
import { PhDs } from 'pages/phds';
import { EventsManage } from 'pages/eventsManage';
import { NewsManage } from 'pages/newsManage';
import { PhdsManage } from 'pages/phdsManage';
import { StaffsManage } from 'pages/staffsManage';
import { ProjectsManage } from 'pages/projectsManage';
import { ResourcesManage } from 'pages/resourceManage';
import { adminLogin, adminLogout, checkIsAdminLogin } from 'utils/request';
import { Form, Input, Checkbox } from '@arco-design/web-react';
import { i18nChangeLanguage } from '@wangeditor/editor'
const FormItem = Form.Item;

const MenuItem = Menu.Item;

i18nChangeLanguage('en')

function App () {
  const [isAdminLogin, setIsAdminLogin] = useState(false);
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  useEffect(() => {
    checkIsAdminLogin().then(res => {
      if (res.code == 0 && res.data.isLogin) {
        setIsAdminLogin(true);
      }
    })
  }, [])


  return (
    <div className="app">
      <Menu mode='horizontal' defaultSelectedKeys={['home']}>
        <MenuItem
          key='logo'
          style={{ padding: 0 }}
          disabled
        >
          <div
            style={{
              // width: 80,
              // height: 30,
              margin: 'auto',
              borderRadius: 2,
              // background: 'var(--color-fill-3)',
              cursor: 'pointer',
            }}
          >
            <Link to='/home'>
              <Avatar size={30} shape='square'>
                Logo
              </Avatar>
            </Link>
          </div>
        </MenuItem>
        <MenuItem key='home'>
          <Link to="/home">Home</Link>
        </MenuItem>
        <MenuItem key='news'>
          <Link to="/news">News</Link>
        </MenuItem>
        <MenuItem key='events'>
          <Link to="/events">Events</Link>
        </MenuItem>
        <MenuItem key='staffs'>
          <Link to="/staffs">Staffs</Link>
        </MenuItem>
        <MenuItem key='phds'>
          <Link to="/phds">PhD Students</Link>
        </MenuItem>
        <MenuItem key='projects'>
          <Link to="/projects">Projects</Link>
        </MenuItem>
        <MenuItem key='resuorces'>
          <Link to="/resuorces">Resources</Link>
        </MenuItem>
        <MenuItem key='contact'>
          <Link to="/contact">Contact</Link>
        </MenuItem>
        <MenuItem key='auth'>
          {isAdminLogin ? <Button type='text' onClick={() => {
            adminLogout().then(res => {
              if (res.code == 0) {
                Message.info('Logout success');
                setIsAdminLogin(false);
              }
            });
          }}>Logout</Button> : <Trigger
            mouseEnterDelay={400}
            mouseLeaveDelay={400}
            trigger='hover'
            popup={() => {
              return (
                <div style={{
                  border: '1px solid var(--color-border-1)',
                  shadow: 'var(--shadow-1)',
                  zIndex: 1000,
                  backgroundColor: 'white',
                  padding: 20,
                }}>
                  <Form style={{ width: 600 }} autoComplete='off'>
                    <FormItem label='Username'>
                      <Input value={username} onChange={setUsername} required placeholder='please enter your username...' />
                    </FormItem>
                    <FormItem label='Post'>
                      <Input value={password} onChange={setPassword} required type='password' placeholder='please enter your password...' />
                    </FormItem>
                    <FormItem wrapperCol={{ offset: 5 }}>
                      <Button type='primary' onClick={() => {
                        adminLogin(username, password).then(res => {
                          if (res.code == 0 && res.data.isLogin) {
                            setIsAdminLogin(true);
                          }
                        })
                      }}>Login</Button>
                    </FormItem>
                  </Form>
                </div>
              )
            }} >
            <Button type='text'>
              Login
            </Button>
          </Trigger>}
        </MenuItem>
      </Menu>

      <Routes>
        <Route path='/home' element={<Home />} />
        <Route path='/contact' element={<Contact />} />
        <Route path='/events' element={<Events />} />
        <Route path='/news' element={<News />} />
        <Route path='/projects' element={<Projects />} />
        <Route path='/resuorces' element={<Resources />} />
        <Route path='/staffs' element={<Staffs />} />
        <Route path='/phds' element={<PhDs />} />
        <Route path='/eventsManage' element={<EventsManage />} />
        <Route path='/newsManage' element={<NewsManage />} />
        <Route path='/phdsManage' element={<PhdsManage />} />
        <Route path='/staffsManage' element={<StaffsManage />} />
        <Route path='/projectsManage' element={<ProjectsManage />} />
        <Route path='/resourcesManage' element={<ResourcesManage />} />
      </Routes>
      <Footer />
    </div>
  );
}

export default App;
