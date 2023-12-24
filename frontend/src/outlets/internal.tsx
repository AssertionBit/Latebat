import { Outlet } from "react-router";
import { Navigator } from "../components";
import style from "./internal.module.css";
import { Suspense } from "react";

export default (): React.JSX.Element => {
    return (
        <main className={style.main}>
            <Navigator/>
            <div className={style.contWrap}>
                <Suspense>
                    <Outlet/>
                </Suspense>
            </div>
        </main>
    );
};
