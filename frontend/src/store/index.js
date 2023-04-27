import { createStore } from 'vuex'


const store = createStore({
    state() {
        return {
            user: {},// 用户信息
        }
    },
    mutations: {
        // 设置用户信息
        SET_USERINFO(state, user) {
            state.user = user
        },
    },
    actions: {
        getinfo({ commit }) {
            return new Promise((resolve, reject) => {

            })
        }
    }
})

export default store