import { Link, useLocation } from "react-router-dom";
import style from "./Navigator.module.css";


export default (): React.JSX.Element => {
    const location = useLocation().pathname;

    return (<>
        <aside className={style.aside}>
            <Link className={`${style.navLink} ${location === "/" ? style.selected : ""}`} to={"/"}>Home</Link>
            <Link className={`${style.navLink} ${location === "/batch" ? style.selected : ""}`} to={"/batch"}>Download</Link>
            <Link className={`${style.navLink} ${location === "/docs" ? style.selected : ""}`} to={"/docs"}>Documents</Link>
            <Link className={`${style.navLink} ${location === "/upload" ? style.selected : ""}`} to={"/upload"}>Upload</Link>
        </aside>
    </>);
};