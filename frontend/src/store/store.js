// store.js

import Vue from 'vue'
import Vuex from 'vuex'
// Vue.use(Vuex)
import { createApp } from 'vue'
import { createStore } from 'vuex'


// const state = {
//     activeRoomId: null // 存储激活的房间信息
// }

// const mutations = {
//     setActiveRoomId(id) {
//         state.activeRoomId = id;
//     }
// }

// const actions = {

// }

// const getters = {
//     activeRoomId(state) {
//         return state.activeRoomId;
//     }

// }

// const store = new Vuex.Store({
//     state,
//     mutations,
//     actions,
//     getters
// })



// export default store

export default new Vuex.Store({
    state: {
        activeRoom: {
            "ID": null,
            "CreatedAt": null,
            "UpdatedAt": null,
            "DeletedAt": null,
            "name": null,
            "description": null
        } // 存储激活的房间信息
    },
    getters: {
        activeRoom(state) {
            return state.activeRoom;
        }
    },
    mutations: {
        setActiveRoom(state, room) {
            // state.activeRoom = room;
            state.activeRoom.ID = room.ID;
            state.activeRoom.CreatedAt = room.CreatedAt;
            state.activeRoom.UpdatedAt = room.UpdatedAt;
            state.activeRoom.DeletedAt = room.DeletedAt;
            state.activeRoom.name = room.name;
            state.activeRoom.description = room.description;

            console.log('state.activeRoom:', state.activeRoom)
        }
    },
    actions: {
        // 不需要在这里定义 fetchRoomInfo 方法，因为这个方法不是从服务器获取数据的
    }
});