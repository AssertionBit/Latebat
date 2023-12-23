import React from "react";

const LoginPage = React.lazy(() => import("./login"));
const SignUpPage = React.lazy(() => import("./signup"));
const AppIndexPage = React.lazy(() => import("./internal-index"));

export {
    AppIndexPage,
    LoginPage,
    SignUpPage,
};
