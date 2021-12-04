export default function ({ store, route, redirect }) {
  if (store.getters.getUser === null && route.path !== "/signin") {
    console.log("redirect")
    return redirect("/signin");
  }
}
