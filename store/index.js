const state = () => ({
  user: null,
  event_ws: null,
  event_org: null,
});

const mutations = {
  SET_USER(state, user) {
    state.user = user;
  },
  connect_event_ws(state, eid) {
    // state.event_ws = new WebSocket(`ws://localhost:8000/api/event/${eid}`)
    state.event_ws = new WebSocket(`ws://commusto-back.herokuapp.com//api/event/${eid}`)
  },
  set_event_org(state, uid) {
    state.event_org = uid
  }
};

const actions = {
  async onAuthStateChangedAction(state, { authUser, claims }) {
    if (!authUser) {
      // authされていない場合
      state.commit("SET_USER", null);

      // リダイレクトの設定
      this.$router.push({
        path: "/signin",
      });
    } else {
      // authされている場合
      const { uid, email } = authUser;
      state.commit("SET_USER", {
        uid,
        email,
      });
    }
  },
};

const getters = {
  getUser(state) {
    return state.user;
  },
  getEventWs(state) {
    return state.event_ws;
  },
  getEventOrg(state) {
    return state.event_org;
  }
};

export default {
  state,
  mutations,
  actions,
  getters,
};
