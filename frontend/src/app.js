import renderNav from "./component/nav/nav.js";
import Home from "./component/page/home/home.js"
import Sign_In from "./component/page/sign_in/sign_in.js"
import Sign_Up from "./component/page/sign_up/sign_up.js"
import Create_Post from "./component/page/create_post/create_post.js";
import Logout from "./component/Log_out.js";
const appDiv = document.getElementById("app");

let usrObj = null;
const router = async() => {
    const routes = [
        { path: "/", pathName: "Home", show: usrObj !== null, view: Home },
        { path: "/sign_in", pathName: "Sign in", show: usrObj === null, view: Sign_In },
        { path: "/sign_up", pathName: "Sign up", show: usrObj === null, view: Sign_Up },
        { path: "/create_post", pathName: "Create post", show: usrObj !== null, view: Create_Post },
        { path: "/logout", pathName: "Log out", show: usrObj !== null, view: Logout },
    ];
    renderNav(routes, appDiv, usrObj)
    const defaultRoute = usrObj === null ? routes[1] : null;

    const matchingRoute = routes.find(
        (route) => route.path === window.location.pathname && route.display
    ) || defaultRoute;

    if (!matchingRoute) {
        appDiv.innerHTML = "<h1>Page not found</h1>";
    } else {
        matchingRoute.view(appDiv)
    }


}

window.addEventListener("popstate", router);


document.addEventListener("DOMContentLoaded", async() => {
    document.body.addEventListener("click", function(e) {
        if (e.target.matches("[data-link]")) {
            e.preventDefault();
            navigateTo(e.target.href);
        }
    });
    router()
});