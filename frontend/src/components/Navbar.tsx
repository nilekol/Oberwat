import React from 'react';
import { Link } from 'react-router-dom'; 
import './../styles/Navbar.css'
import Name from './Name';
import SearchBar from './SearchBar';

const Navbar: React.FC = () => {
  return (
    <nav className="navbar">
      <div className="navbar-container">
        <Link to="/" className="navbar-logo">
          <Name />
        </Link>
        <div className="navbar-tabs">
          <Link to="/tab1" className="navbar-tab">Heroes</Link>
          <Link to="/tab2" className="navbar-tab">Maps</Link>
          <Link to="/tab3" className="navbar-tab">About Me</Link>
        </div>
        <div className="search-bar-container">
          <SearchBar onSearch={(query) => console.log(query)} />
        </div>
      </div>
    </nav>
  );
}

export default Navbar;
