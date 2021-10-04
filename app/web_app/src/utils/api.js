import axios from 'axios';


const apiInstance = axios.create({
    baseURL: 'http://178.79.148.75',
    header: {
        //TODO(sam): Any headers we might need we can put here
    },
    auth: {
        username: 'admin',
        password: 'blipblop'
    }
});

const api = {
    fetchJobs: () => {
        apiInstance({
            'method': 'GET',
            'url': '/jobs',
        })
    }
}

export default api;