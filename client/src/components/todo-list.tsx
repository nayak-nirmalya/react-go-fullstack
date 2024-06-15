import { useState } from "react";
import { Flex, Spinner, Stack, Text } from "@chakra-ui/react";

import { ToDoItem } from "./todo-item";

import { ToDo } from "../types";

export function ToDoList() {
  const [isLoading, setIsLoading] = useState(true);

  const todos: ToDo[] = [{ id: 1, body: "Learn React", completed: true }];

  return (
    <>
      <Text
        fontSize={"4xl"}
        textTransform={"uppercase"}
        fontWeight={"bold"}
        textAlign={"center"}
        my={2}
      >
        {`Today's Task${todos?.length === 1 ? "" : "s"}`}
      </Text>
      {isLoading && (
        <Flex justifyContent={"center"} my={4}>
          <Spinner size={"xl"} />
        </Flex>
      )}
      {!isLoading && todos?.length === 0 && (
        <Stack alignItems={"center"} gap="3">
          <Text fontSize={"xl"} textAlign={"center"} color={"gray.500"}>
            All Tasks Completed! 🎉
          </Text>
          <img src="/go.png" alt="Go logo" width={70} height={70} />
        </Stack>
      )}
      <Stack gap={3}>
        {todos?.map((todo) => (
          <ToDoItem key={todo.id} todo={todo} />
        ))}
      </Stack>
    </>
  );
}
