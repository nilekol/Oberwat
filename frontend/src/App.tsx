import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import BattleTag from './pages/BattleTag'

import Home from './pages/Home'

function App() {
    return (
        <Router>
            <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/:battletag" element={<BattleTag />} />

            </Routes>
        </Router>
    )
  }
  
  export default App