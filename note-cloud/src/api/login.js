import axios from 'axios'

export function fetch(url, params) {
    return new Promise((resolve, reject) => {
        axios.get(url, params)
             .then(response => {
                 resolve(response.data)
             })
             .catch((error) => {
                 console.log(error)
                 reject(error)
             })
    })
}
export function postFetch(url, params,config) {
    return new Promise((resolve, reject) => {
        axios.post(url, params,config)
             .then(response => {
                 resolve(response.data)
             })
             .catch((error) => {
                 console.log(error)
                 reject(error)
             })
    })
}

export default {
    postUserRegister(params) {
        return postFetch("api/register",params)
    },
    postUserLogin(params) {
        return postFetch("api/signin",params)
    }
}