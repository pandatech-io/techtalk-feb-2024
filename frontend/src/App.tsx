import {Container, Button} from 'react-bootstrap'
import { useGetIds } from './api/generated/services/doc/default/default';
import { GetIds200 } from './api/generated/models/doc';


function App() {
  const {data, refetch}=  useGetIds()
  return (
    <Container className="p-3">
      <Container className="pt-5 mb-4">
        <h1 className='text-center mb-3'>ID Generator</h1>
        <h3 className='text-center'>{(data as GetIds200)?.data}</h3>
      </Container>
      <Container className='d-flex justify-content-center'>
        <Button className="mx-1" onClick={() => refetch()}>
          Generate
        </Button>
      </Container>
    </Container>
  );
}

export default App;
