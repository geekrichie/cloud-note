import api from "../api/login"
const state ={
    isLogin : false,
    username : '',
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
             commit("login", true)
             let data = res.data
             if (data != null) {
                 commit("setUserName", data.username)
             }
         }
     }
}

export default  {
    namespaced:true,
    state,
    mutations,
    actions
}
