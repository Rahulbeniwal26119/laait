import React from "react";
import Box from "@mui/material/Box";
import { Navbar, Container, Nav } from "react-bootstrap";

export default class NavBar extends React.Component {
  render() {
    return (
      <Navbar collapseOnSelect 
       fixed="top" bg="dark" variant="dark">
        <Container>
          <Navbar.Brand href="#home">Laait</Navbar.Brand>
          <Nav navbarScroll className="me-auto">
            <Nav.Link href="#home">Home</Nav.Link>
            <Nav.Link href="#features">Features</Nav.Link>
            <Nav.Link href="#download">Download</Nav.Link>
          </Nav>
        </Container>
      </Navbar>
    );
  }
}
