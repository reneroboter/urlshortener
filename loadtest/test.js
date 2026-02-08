import http from 'k6/http';
import { SharedArray } from 'k6/data';

// CSV laden (einmal)
const data = new SharedArray('urls', function () {
    const file = open('top-1m.csv');
    const lines = file.split('\n');

    // CSV: rank,domain
    return lines.map(line => {
        const parts = line.split(',');
        return parts[1] ? parts[1].trim() : null;
    }).filter(Boolean);
});

export const options = {
    vus: 200,
    duration: '1m',
};

export default function () {
    // zufällige URL auswählen
    const url = data[Math.floor(Math.random() * data.length)];

    const payload = JSON.stringify({
        url: `http://${url}`,
    });

    const res = http.post(
        'http://127.0.0.1:8888/shorten',
        payload,
        { headers: { 'Content-Type': 'application/json' } }
    );
}
