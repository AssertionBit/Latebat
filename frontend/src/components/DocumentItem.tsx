import { useEffect, useState } from "react";
import style from "./DocumentItem.module.css";
import Badge from "./Badge";

export interface IDocumentItemProps {
    id: number,
    name: string,
    type: string,
    status: string,
}

export default (props: IDocumentItemProps, key: number): React.JSX.Element => {
    const [clicked, setClicked] = useState<boolean>(false);

    useEffect(() => {
        const dialog = document.querySelector("dialog");

        if(clicked) {
            dialog?.showModal();
        } else {
            dialog?.close();
        }
    }, [clicked]);

    return (
        <>
            <div key={`${key}`} onClick={() => setClicked(!clicked)}
                 className={style.wrapper}>
                <h3 className={style.itemName}>{props.name}#{props.id}</h3>
                <Badge text={props.type}/><Badge text={props.status}/>
                <button className={style.itemOpener}>Open info</button>
            </div>

            {clicked && <dialog className={style.dialog}>
                <form method="dialog">
                    <h2>He</h2>
                    <button 
                        className={style.formCloseBtn}
                        onClick={() => setClicked(!clicked)}>
                        Close
                    </button>
                </form>
            </dialog> }
        </>
    );
};