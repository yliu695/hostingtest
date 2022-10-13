import axios from 'axios';

axios.defaults.withCredentials = true

const instance = axios.create({
    baseURL: 'http://localhost:3000/api',
    withCredentials: true,
});

export async function eventList () {
    try {
        let res = await instance.get("/events?order=event_time desc")

        if (res.status != 200) {
            return {
                code: 1,
                msg: 'request failed with status code ' + res.status
            }
        }

        return {
            code: 0,
            data: res.data
        }
    }
    catch (err) {
        return {
            code: 1,
            msg: err,
        }
    }
}

export async function addEvent (title, content, tags, cover, event_time) {
    try {
        let res = await instance.post("/events", {
            "title": title,
            "content": content,
            "tags": tags.join('|'),
            "cover": cover,
            "event_time": event_time | 0,
        })

        if (res.status != 200) {
            return {
                code: 1,
                msg: 'request failed with status code ' + res.status
            }
        }

        return {
            code: 0,
            data: res.data
        }
    }
    catch (err) {
        return {
            code: 1,
            msg: err,
        }
    }
}

export async function editEvent (eventId, title, content, tags, cover, event_time) {
    try {
        let res = await instance.put(`/events/${eventId}`, {
            "title": title,
            "content": content,
            "tags": tags.join('|'),
            "cover": cover,
            "event_time": event_time | 0,
        })

        if (res.status != 200) {
            return {
                code: 1,
                msg: 'request failed with status code ' + res.status
            }
        }

        return {
            code: 0,
            data: res.data
        }
    }
    catch (err) {
        return {
            code: 1,
            msg: err,
        }
    }
}

export async function deleteEvent (eventId) {
    try {
        let res = await instance.delete(`/events/${eventId}`)

        if (res.status != 200) {
            return {
                code: 1,
                msg: 'request failed with status code ' + res.status
            }
        }

        return {
            code: 0,
            data: res.data
        }
    }
    catch (err) {
        return {
            code: 1,
            msg: err,
        }
    }
}

export async function newsList () {
    try {
        let res = await instance.get("/news?order=create_time desc")

        if (res.status != 200) {
            return {
                code: 1,
                msg: 'request failed with status code ' + res.status
            }
        }

        return {
            code: 0,
            data: res.data
        }
    }
    catch (err) {
        return {
            code: 1,
            msg: err,
        }
    }
}

export async function addNews (title, content, tags, cover) {
    try {
        let res = await instance.post("/news", {
            "title": title,
            "content": content,
            "tags": tags.join('|'),
            "cover": cover,
        })

        if (res.status != 200) {
            return {
                code: 1,
                msg: 'request failed with status code ' + res.status
            }
        }

        return {
            code: 0,
            data: res.data
        }
    }
    catch (err) {
        return {
            code: 1,
            msg: err,
        }
    }
}

export async function editNews (newsId, title, content, tags, cover) {
    try {
        let res = await instance.put(`/news/${newsId}`, {
            "title": title,
            "content": content,
            "tags": tags.join('|'),
            "cover": cover,
        })

        if (res.status != 200) {
            return {
                code: 1,
                msg: 'request failed with status code ' + res.status
            }
        }

        return {
            code: 0,
            data: res.data
        }
    }
    catch (err) {
        return {
            code: 1,
            msg: err,
        }
    }
}

export async function deleteNews (newsId) {
    try {
        let res = await instance.delete(`/news/${newsId}`)

        if (res.status != 200) {
            return {
                code: 1,
                msg: 'request failed with status code ' + res.status
            }
        }

        return {
            code: 0,
            data: res.data
        }
    }
    catch (err) {
        return {
            code: 1,
            msg: err,
        }
    }
}

export async function phdList () {
    try {
        let res = await instance.get("/phds")

        if (res.status != 200) {
            return {
                code: 1,
                msg: 'request failed with status code ' + res.status
            }
        }

        return {
            code: 0,
            data: res.data
        }
    }
    catch (err) {
        return {
            code: 1,
            msg: err,
        }
    }
}

export async function addPhd (name, job, intro, avatar) {
    try {
        let res = await instance.post("/phds", {
            name: name,
            job: job,
            intro: intro,
            avatar: avatar,
        })

        if (res.status != 200) {
            return {
                code: 1,
                msg: 'request failed with status code ' + res.status
            }
        }

        return {
            code: 0,
            data: res.data
        }
    }
    catch (err) {
        return {
            code: 1,
            msg: err,
        }
    }
}

export async function editPhd (phdId, name, job, intro, avatar) {
    try {
        let res = await instance.put(`/phds/${phdId}`, {
            name: name,
            job: job,
            intro: intro,
            avatar: avatar,
        })

        if (res.status != 200) {
            return {
                code: 1,
                msg: 'request failed with status code ' + res.status
            }
        }

        return {
            code: 0,
            data: res.data
        }
    }
    catch (err) {
        return {
            code: 1,
            msg: err,
        }
    }
}

export async function deletePhd (phdId) {
    try {
        let res = await instance.delete(`/phds/${phdId}`)

        if (res.status != 200) {
            return {
                code: 1,
                msg: 'request failed with status code ' + res.status
            }
        }

        return {
            code: 0,
            data: res.data
        }
    }
    catch (err) {
        return {
            code: 1,
            msg: err,
        }
    }
}

export async function staffList () {
    try {
        let res = await instance.get("/staffs")

        if (res.status != 200) {
            return {
                code: 1,
                msg: 'request failed with status code ' + res.status
            }
        }

        return {
            code: 0,
            data: res.data
        }
    }
    catch (err) {
        return {
            code: 1,
            msg: err,
        }
    }
}

export async function addStaff (name, job, intro, avatar) {
    try {
        let res = await instance.post("/staffs", {
            name: name,
            job: job,
            intro: intro,
            avatar: avatar,
        })

        if (res.status != 200) {
            return {
                code: 1,
                msg: 'request failed with status code ' + res.status
            }
        }

        return {
            code: 0,
            data: res.data
        }
    }
    catch (err) {
        return {
            code: 1,
            msg: err,
        }
    }
}

export async function editStaff (staffId, name, job, intro, avatar) {
    try {
        let res = await instance.put(`/staffs/${staffId}`, {
            name: name,
            job: job,
            intro: intro,
            avatar: avatar,
        })

        if (res.status != 200) {
            return {
                code: 1,
                msg: 'request failed with status code ' + res.status
            }
        }

        return {
            code: 0,
            data: res.data
        }
    }
    catch (err) {
        return {
            code: 1,
            msg: err,
        }
    }
}

export async function deleteStaff (staffId) {
    try {
        let res = await instance.delete(`/staffs/${staffId}`)

        if (res.status != 200) {
            return {
                code: 1,
                msg: 'request failed with status code ' + res.status
            }
        }

        return {
            code: 0,
            data: res.data
        }
    }
    catch (err) {
        return {
            code: 1,
            msg: err,
        }
    }
}

export async function projectList () {
    try {
        let res = await instance.get("/projects")

        if (res.status != 200) {
            return {
                code: 1,
                msg: 'request failed with status code ' + res.status
            }
        }

        return {
            code: 0,
            data: res.data
        }
    }
    catch (err) {
        return {
            code: 1,
            msg: err,
        }
    }
}

export async function addProject (name, intro) {
    try {
        let res = await instance.post("/projects", {
            name: name,
            intro: intro,
        })

        if (res.status != 200) {
            return {
                code: 1,
                msg: 'request failed with status code ' + res.status
            }
        }

        return {
            code: 0,
            data: res.data
        }
    }
    catch (err) {
        return {
            code: 1,
            msg: err,
        }
    }
}

export async function editProject (projectId, name, intro) {
    try {
        let res = await instance.put(`/projects/${projectId}`, {
            name: name,
            intro: intro,
        })

        if (res.status != 200) {
            return {
                code: 1,
                msg: 'request failed with status code ' + res.status
            }
        }

        return {
            code: 0,
            data: res.data
        }
    }
    catch (err) {
        return {
            code: 1,
            msg: err,
        }
    }
}

export async function deleteProject (projectId) {
    try {
        let res = await instance.delete(`/projects/${projectId}`)

        if (res.status != 200) {
            return {
                code: 1,
                msg: 'request failed with status code ' + res.status
            }
        }

        return {
            code: 0,
            data: res.data
        }
    }
    catch (err) {
        return {
            code: 1,
            msg: err,
        }
    }
}

export async function resourceList () {
    try {
        let res = await instance.get("/resources")

        if (res.status != 200) {
            return {
                code: 1,
                msg: 'request failed with status code ' + res.status
            }
        }

        return {
            code: 0,
            data: res.data
        }
    }
    catch (err) {
        return {
            code: 1,
            msg: err,
        }
    }
}

export async function addResource (name, intro, link) {
    try {
        let res = await instance.post("/resources", {
            name: name,
            intro: intro,
            link: link,
        })

        if (res.status != 200) {
            return {
                code: 1,
                msg: 'request failed with status code ' + res.status
            }
        }

        return {
            code: 0,
            data: res.data
        }
    }
    catch (err) {
        return {
            code: 1,
            msg: err,
        }
    }
}

export async function editResource (resourceId, name, intro, link) {
    try {
        let res = await instance.put(`/resources/${resourceId}`, {
            name: name,
            intro: intro,
            link: link,
        })

        if (res.status != 200) {
            return {
                code: 1,
                msg: 'request failed with status code ' + res.status
            }
        }

        return {
            code: 0,
            data: res.data
        }
    }
    catch (err) {
        return {
            code: 1,
            msg: err,
        }
    }
}

export async function deleteResource (resourceId) {
    try {
        let res = await instance.delete(`/resources/${resourceId}`)

        if (res.status != 200) {
            return {
                code: 1,
                msg: 'request failed with status code ' + res.status
            }
        }

        return {
            code: 0,
            data: res.data
        }
    }
    catch (err) {
        return {
            code: 1,
            msg: err,
        }
    }
}

export async function checkIsAdminLogin () {
    try {
        let res = await instance.get(`/isAdminLogin`)

        if (res.status != 200) {
            return {
                code: 1,
                msg: 'request failed with status code ' + res.status
            }
        }

        return {
            code: 0,
            data: res.data
        }
    }
    catch (err) {
        return {
            code: 1,
            msg: err,
        }
    }
}

export async function adminLogin (username, password) {
    try {
        let res = await instance.post(`/adminLogin?username=${username}&password=${password}`)

        if (res.status != 200) {
            return {
                code: 1,
                msg: 'request failed with status code ' + res.status
            }
        }

        return {
            code: 0,
            data: res.data
        }
    }
    catch (err) {
        return {
            code: 1,
            msg: err,
        }
    }
}

export async function adminLogout () {
    try {
        let res = await instance.post(`/adminLogout`)

        if (res.status != 200) {
            return {
                code: 1,
                msg: 'request failed with status code ' + res.status
            }
        }

        return {
            code: 0,
            data: res.data
        }
    }
    catch (err) {
        return {
            code: 1,
            msg: err,
        }
    }
}

export async function notifyContact (name, email, feedback) {
    try {
        let res = await instance.post(`/notifyContact?name=${name}&email=${email}&feedback=${feedback}`)

        if (res.status != 200) {
            return {
                code: 1,
                msg: 'request failed with status code ' + res.status
            }
        }

        return {
            code: 0,
            data: res.data
        }
    }
    catch (err) {
        return {
            code: 1,
            msg: err,
        }
    }
}