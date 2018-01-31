import ReactDOM from 'react-dom';
import React from 'react';

const Avatar = props => {
    console.log(props)
    return <img className="img-fluid" src={`/${props.user.avatar}`}  alt=""/>
}

export { Avatar }