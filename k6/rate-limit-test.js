import http from 'k6/http';
import { check } from 'k6';

export const options = {
    scenarios: {
        rate_limit_test: {
            executor: 'constant-arrival-rate',
            rate: 1200,       // RPS
            timeUnit: '1s',
            duration: '30s',
            preAllocatedVUs: 50,
            maxVUs: 200,
        },
    },
};

const BASE_URL = 'http://127.0.0.1:3000';
const API_KEY = 'd930ffe9999d518914a28ae89fd991f0f24b8de7ef9ef01e0d2a4343dadb69c3';

export default function () {
    const res = http.post(
        `${BASE_URL}/api/logs`,
        null,
        {
            headers: {
                'X-API-Key': API_KEY,
            },
        }
    );

    check(res, {
        '200 OK': (r) => r.status === 200,
        '429 Too Many Requests': (r) => r.status === 429,
    });
}
