import api from "../api/login"

const user = JSON.parse(localStorage.getItem('user'));

const state = {
    isLogin : user? true:false,
    username : user?user.username:'',
}

const mutations = {
      login(state, isLogin) {         
           state.isLogin = isLogin
      },
      setUserName(state, name) {
          state.username = name
      }
}

const actions = {
     async userLogin({commit},userInfo) {
         const res = await api.postUserLogin(userInfo)
         if (res.code == 0) {
             localStorage.setItem('user', JSON.stringify(res.data));
             commit("login", true)
             commit("setUserName", res.data.username)
         }else {
             commit("login", false)
             commit("setUserName", null)
         }
     }
}

export default  {
    namespaced:true,
    state,
    mutations,
    actions
}
