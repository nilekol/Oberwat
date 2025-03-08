import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import BattleTag from './pages/BattleTag'

import Home from './pages/Home'
import Heroes from './pages/Heroes'
import Maps from './pages/Maps'

function App() {
    return (
        <Router>
            <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/:battletag" element={<BattleTag />} />
            <Route path="/Heroes" element={<Heroes />} />
            <Route path="/Maps" element={<Maps />} />
            </Routes>
        </Router>
    )
  }
  
  export default App