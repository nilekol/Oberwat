import React, { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import './../styles/Navbar.css';
import Name from './Name';
import SearchBar from './SearchBar';
import axios from 'axios';

const Navbar: React.FC = () => {
  const [battletag, setBattletag] = useState('');
  const navigate = useNavigate();

  const handleSearch = async (query: string) => {
    if (query.trim() === "") {
      return;
    }

    query = query.replace("#", "-");

    try {
      const response = await axios.get(`http://localhost:8080/api/players/${query}/summary`);
      console.log(response.data);
      // Navigate and pass the data to the new page
      
      navigate(`/${query}`, { state: { playerData: response.data } });
      
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <nav className="navbar">
      <div className="navbar-container">
        <Link to="/" className="navbar-logo">
          <Name />
        </Link>
        <div className="navbar-tabs">
          <Link to="/Heroes" className="navbar-tab">Heroes</Link>
          <Link to="/Maps" className="navbar-tab">Maps</Link>
          <Link to="/tab3" className="navbar-tab">About Me</Link>
        </div>
        <div className="search-bar-container">
          <SearchBar onSearch={(query) => handleSearch(query)} />
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
