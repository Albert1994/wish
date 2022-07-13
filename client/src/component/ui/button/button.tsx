import React, {FC} from 'react';

interface ButtonProps {
    width?: string;
    height?: string;
    variant: ButtonVariant;
    children?: React.ReactNode;
}

export enum ButtonVariant {
    primary = "primary",
    outlined = 'outlined',
}

const Button: FC<ButtonProps> = ({width, height, children, variant}) => {
    return (
        <button
            className={"w-52 opacity-85 h-9 text-lg p-1 rounded-lg shadow-md bg-gradient-to-r from-green-400 to-blue-500 hover:from-pink-500 hover:to-yellow-500"}
            style={{height,
            width,
            border: variant === ButtonVariant.outlined ? '' : 'none',
            background: variant === ButtonVariant.primary ? '' : '',
            color: 'black'}}>
            {children}
        </button>
    );
};

export default Button;
