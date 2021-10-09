import React from "react";
import { Text } from "react-native";
import { styles } from "./styles";

export default function Tag({tag}) {
    return (
        <Text style={[styles.tag, tag.matched ? styles.tagMatch : styles.tagNoMatch]}>{tag.name}</Text>
    );
}