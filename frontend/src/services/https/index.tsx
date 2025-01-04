import { AppointmentInterface } from "../../interface/IAppointment";
import { SignInInterface } from "../../interface/SignIn";
import axios from "axios";
const apiUrl = "http://localhost:8000";
const Authorization = localStorage.getItem("token");
const Bearer = localStorage.getItem("token_type");
const requestOptions = {
  headers: {
    "Content-Type": "application/json",
    Authorization: `${Bearer} ${Authorization}`,
  },
};
async function SignIn(data: SignInInterface) {
  return await axios
    .post(`${apiUrl}/signin`, data, requestOptions)
    .then((res) => res)
    .catch((e) => e.response);
}

async function GetDepartment() {
  return await axios
    .get(`${apiUrl}/departments`,requestOptions)
    .then((res) => res)
    .catch((e) => e.response);
}

async function GetRoom() {
  return await axios
    .get(`${apiUrl}/rooms`,requestOptions)
    .then((res) => res)
    .catch((e) => e.response);
}

async function GetTime() {
  return await axios
    .get(`${apiUrl}/times`,requestOptions)
    .then((res) => res)
    .catch((e) => e.response);
}

async function CreateAppointment(data: AppointmentInterface) {
  return await axios
    .post(`${apiUrl}/create`, data, requestOptions)
    .then((res) => res)
    .catch((e) => e.response);
}

async function GetAppointment() {
  return await axios
    .get(`${apiUrl}/appointments`,requestOptions)
    .then((res) => res)
    .catch((e) => e.response);
}
export {
  SignIn,
  GetDepartment,
  GetRoom,
  GetTime,
  CreateAppointment,
  GetAppointment
  
};