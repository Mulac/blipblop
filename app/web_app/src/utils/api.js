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

export default api = {
    fetchJobs: () =>
        apiInstance({
            'method': 'GET',
            'url': '/jobs',
            transformResponse: [function (data) {
                data = JSON.parse(data);

                // We can define the name for each attribute here, so if they change from the response in the future
                // we do not need to go digging through the code and change all the individual calls to attributes
                jobs = [];
                data.forEach(job => {
                    jobs.push({
                        Company: job.Company,
                        CompanyImage: job.CompanyImage,
                        Location: job.Location,
                        Title: job.Title,
                        Description: job.Description
                    });
                });

                return jobs
            }]
        })
};