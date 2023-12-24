import styles from "./Badge.module.css";

export interface IBadgeProps {
    text: string;
}

export default (props: IBadgeProps): React.JSX.Element => {
    return (
        <span className={styles.badge}>
            {props.text}
        </span>
    );
};