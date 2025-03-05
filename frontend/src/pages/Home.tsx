import SearchBar from '../components/SearchBar';
import Name from '../components/Name';

function Home() {
    return (
        <div>
            <SearchBar onSearch={(query) => console.log(query)} />
            <Name/>
            <h1>Home</h1>
        </div>
    )
}

export default Home