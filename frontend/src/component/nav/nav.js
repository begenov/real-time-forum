export default function renderNav(routes, appDiv, usrObj) {
    const navContainer = document.createElement("div");
    navContainer.classList.add("navigation");
    routes.forEach(route => {
        if (!route.show) {
            return
        }
        const link = document.createElement("a");
        link.href = route.path;
        link.textContent = route.pathName;
        navContainer.appendChild(link);
    });
    if (usrObj) {
        const link = document.createElement("a");
        link.href = "/profile";
        link.textContent = usrObj;
        navContainer.appendChild(link);

    }
    appDiv.appendChild(navContainer);
}