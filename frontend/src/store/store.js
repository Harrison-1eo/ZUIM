// store.js

import Vuex from 'vuex'

const store = new Vuex.Store({
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
        getActiveRoom(state) {
            if (!state.activeRoom.ID) {
                console.error('activeRoom is null');
                return null;
            }
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
        },
        clearActiveRoom(state) {
            state.activeRoom.ID = null;
            state.activeRoom.CreatedAt = null;
            state.activeRoom.UpdatedAt = null;
            state.activeRoom.DeletedAt = null;
            state.activeRoom.name = null;
            state.activeRoom.description = null;
        },
    },
    actions: {
        // 不需要在这里定义 fetchRoomInfo 方法，因为这个方法不是从服务器获取数据的
    },
});

export default store