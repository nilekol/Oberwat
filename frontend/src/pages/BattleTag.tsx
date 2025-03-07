import { useParams, useLocation } from 'react-router-dom';

function BattleTag() {
    // Get battletag from URL parameters
    const { battletag } = useParams<{ battletag: string }>();

    // Get player data passed through state from the Navbar component
    const location = useLocation();
    const playerData = location.state?.playerData;

    return (
        <div>
            <h1>{battletag}</h1>

            {/* Check if player data is available */}
            {playerData ? (
                <div>
                    <h2>Player Data:</h2>
                    <pre>{JSON.stringify(playerData, null, 2)}</pre>
                </div>
            ) : (
                <p>No player data found</p>
            )}
        </div>
    );
}

export default BattleTag;
