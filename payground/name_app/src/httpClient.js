// import *as storage from "./storage";
import axios from "axios";
import * as storage from "./storage.js";
import router from "./router.js";

function postRequest(method, url, data) {
  return axios.post("http://localhost:8080" + url, JSON.stringify(data), {
    headers: {
      Authorization: storage.getToken(),
    },
  });
}

function getRequest(method, url) {
  return axios.get("http://localhost:8080" + url, {
    headers: {
      Authorization: storage.getToken(),
    },
  });
}

export function Post(url, sendData) {
  return postRequest("POST", url, sendData);
}

export function Get(url) {
  const request = getRequest("GET", url);

  return request.catch((e) => {
    if (e.response.status === 401) {
      router.push("/");
    }
  });
}
