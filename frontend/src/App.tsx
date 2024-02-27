import {Container, Button} from 'react-bootstrap'


function App() {
  return (
    <Container className="p-3">
      <Container className="pt-5 mb-4">
        <h1 className='text-center mb-3'>ID Generator</h1>
        <h3 className='text-center'>generated_id</h3>
      </Container>
      <Container className='d-flex justify-content-center'>
        <Button className="mx-1" onClick={() => {}}>
          Generate
        </Button>
      </Container>
    </Container>
  );
}

export default App;
