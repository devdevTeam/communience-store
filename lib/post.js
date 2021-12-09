import axios from 'axios'
async function post(endpoint, params) {
    // let uri = "http://localhost:8000/api" + endpoint 
    let uri = "http://localhost:8000/api" + endpoint 
    let res = await axios.post(uri, params)
    return res
}

export default post