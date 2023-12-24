import React from "react";

const LoginPage = React.lazy(() => import("./login"));
const SignUpPage = React.lazy(() => import("./signup"));
const AppIndexPage = React.lazy(() => import("./internal-index"));
const AppBatchPage = React.lazy(() => import("./download"));

export {
    AppIndexPage,
    AppBatchPage,
    LoginPage,
    SignUpPage,
};
