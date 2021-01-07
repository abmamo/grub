import { Route, Switch, Link } from 'react-router-dom';
import styled from 'styled-components';
import HomeIcon from '@material-ui/icons/Home';

const AppWrapper = styled.section`
  padding: 4em;
  background: white;
  margin: 0 auto;
  max-height: 100vh;
  max-width: 100vw;
  height: 100%;
  width: 100%;
  text-align: center;
`;

const AppTitle = styled.h1`
  font-size: 5rem;
  border-bottom: 2px solid #1DA76F;
`;

const Input = styled.input`
  font-size: 1rem;
  display: block;
  border-radius: 3px;
  padding: 1.5em;
  margin: 0 auto;
  margin-top: 1rem;
  margin-bottom: 1rem;
  color: ${props => props.inputColor || "palevioletred"};
  background: transparent;
  border: 3px solid #1DA76F;
  width: 250px;
  max-width: 50%;
`;

const Button = styled.button`
  font-size: 1.5rem;
  display: block;
  border-radius: 3px;
  padding: 1em;
  margin: 0 auto;
  margin-top: 1rem;
  margin-bottom: 1rem;
  background: transparent;
  color: #1DA76F;
  border: 3px solid #1DA76F;
  text-align: center;
  transition: 0.3s ease;
  max-width: 25%;
  width: 250px;
  text-decoration: none;
  
  &:hover {
    color: #F8F8F8;
    background: #1DA76F;
  }
`

function Home() {
  return (
    <div>
      <Button as={Link} to="/create">
        create
      </Button>
      <Button as={Link} to="/join">
        join
      </Button>
      <Button as={Link} to="/manage">
        manage
      </Button>
    </div>
  )
}

function Create() {
  return (
    <div>
      <Input defaultValue="menu link" type="text" />
      <Input defaultValue="expires" type="text" />
      <Input defaultValue="name" type="text" />
      <Button as={Link} to="/create">
        create
      </Button>
      <Button as={Link} to="/">
        <HomeIcon />
      </Button>
    </div>

  )
}

function Join() {
  return (
    <div>
      <Input defaultValue="cart ID" type="text" />
      <Input defaultValue="name" type="text" />
      <Input defaultValue="items" type="text" />
      <Input defaultValue="notes" type="text" />
      <Button as={Link} to="/create">
        join
      </Button>
      <Button as={Link} to="/">
        <HomeIcon />
      </Button>
    </div>
  )
}

function Manage() {
  return (
    <div>
      <Input defaultValue="cart ID" type="text" />
      <Input defaultValue="edit expiration" type="text" />
      <Input defaultValue="edit menu" type="text" />
      <Button as={Link} to="/home">
        update
      </Button>
      <Button as={Link} to="/">
        <HomeIcon />
      </Button>
    </div>
  )
}

function Header() {
  return (
    <AppTitle>
      <h1>grub</h1>
    </AppTitle>
  );
}

function Main() {
  return (
    <main>
      <Header />
      <Switch>
        <Route exact path='/' component={Home} />
        <Route path='/create' component={Create} />
        <Route path='/join' component={Join} />
        <Route path='/manage' component={Manage} />
      </Switch>
    </main>
  );
}

function App() {
  return (
    <div>
      <AppWrapper>
        <Main />
      </AppWrapper>
    </div>
  );
}

export default App;
