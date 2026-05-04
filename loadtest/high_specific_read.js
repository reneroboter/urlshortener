import http from 'k6/http';
import { check } from 'k6';

export const options = {
    scenarios: {
        read_test: {
            executor: 'constant-arrival-rate',
            rate: 16000,
            timeUnit: '1s',
            duration: '1m',
            preAllocatedVUs: 100,
            maxVUs: 300,
        },
    },
};

export default function () {
    const res = http.get('http://127.0.0.1:64419/738ddf35b3a85a7a6ba7b232bd3d5f1e4d284ad1', {
        redirects: 0,
        timeout: '2s',
    });

    check(res, {
        'is redirect': r =>
            r.status === 301 ||
            r.status === 302 ||
            r.status === 307 ||
            r.status === 308,
    });
}