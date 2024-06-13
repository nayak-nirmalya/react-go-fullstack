import { Container, Stack } from "@chakra-ui/react";

import { Navbar } from "./components/navbar";

export default function App() {
  return (
    <Stack h="100vh">
      <Navbar />
      <Container>
        {/* ToDoForm */}
        {/* ToDoList */}
      </Container>
    </Stack>
  );
}
