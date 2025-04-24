import axios from "axios";

const api = import.meta.env.VITE_API;

export async function shortenURL(data) {
    const res = await axios.post(`${api}/shortener`, data)
    return res.data;
}

export async function getUrls(data) {
    const res = await axios.get(`${api}/urls/${data}`,)
    return res.data;
}

export async function getAccessLog(shortCode) {
    const res = await axios.get(`${api}/${shortCode}/access_log`)
    return res.data;
}

export async function redirection(data) {
    const res = await axios.post(`${api}/redirect`, data)
    return res.data;
}

export async function getIpInfo() {
    const res = await axios.get("https://ipinfo.io/json?token=ed6ebd504c799d")
    return res.data
}
