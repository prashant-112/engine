import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import { AppBar, Toolbar, Typography, Container, Box } from '@mui/material';
import Search from './components/Search';
import Upload from './components/Upload';

function App() {
  return (
    <Router>
      <Box sx={{ flexGrow: 1 }}>
        <AppBar position="static">
          <Toolbar>
            <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
              Search Engine
            </Typography>
            <Box sx={{ display: 'flex', gap: 2 }}>
              <Link to="/" style={{ color: 'white', textDecoration: 'none' }}>
                Search
              </Link>
              <Link to="/upload" style={{ color: 'white', textDecoration: 'none' }}>
                Upload
              </Link>
            </Box>
          </Toolbar>
        </AppBar>
        <Container>
          <Routes>
            <Route path="/" element={<Search />} />
            <Route path="/upload" element={<Upload />} />
          </Routes>
        </Container>
      </Box>
    </Router>
  );
}

export default App;
