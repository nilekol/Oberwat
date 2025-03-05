import './../styles/Name.css';
import logo from './../assets/logo.png';  // Replace with your actual logo path

function Name() {
    return (
        <div className="name-container">
            <h1>
                <span className="ober">Ober</span>
                <span className="wat">wat</span>
            </h1>
            <img src={logo} alt="Oberwat Logo" />
        </div>
    )
}

export default Name;
