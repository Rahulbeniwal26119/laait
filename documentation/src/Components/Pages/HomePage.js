import React from 'react'
import styled from 'styled-components'
import {Button} from 'react-bootstrap'

const HomePage = () => {
    return (
        <StyledHomePage>
            <Heading>Home Page</Heading>
        </StyledHomePage>
    )
}

const StyledHomePage = styled.div`
    min-height: 100vh;
    width: 100vw;
    background-color: #282c34;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
`;

const Heading = styled.h1`
    font-size: clamp(3rem, 5vw, 7vw);
    color: #eee;
    font-weight: 300;
    margin: 0;
    padding: 0;
    user-select: none;
`;

export default HomePage