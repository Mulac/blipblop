// Card stylesheet

import { StyleSheet } from "react-native";

export const styles = StyleSheet.create({
    container: {
        width: '100%',
        height: '90%',
        // backgroundColor: '#A7C7E7',
        backgroundColor: "#fff",

        borderTopLeftRadius: 30,
        borderTopRightRadius: 30,
        position: 'absolute',
        flex: 1,
        bottom: 0,
    },
  
    companyDetails: {
        flexDirection: 'row',
        alignItems: 'center'
    },  
    companyName: {
        fontSize: 15,
        fontWeight: 'bold',
        color: "#3d3d3d"
    },
    positionName: {
        marginTop: 10,
        fontSize: 20,
        fontWeight: 'bold',
        color: 'black',
        color: "#3d3d3d"
    },
    companyImage: {
        width: 70,
        height: 50,
        marginRight: 10,

        // TODO(sam): Need to find a way to stop this flickering on re-rending
        resizeMode: 'contain'
    },

    locationDetails: {
        flexDirection: 'row',
        alignItems: 'center',
        marginTop: 20
    },
    locationPin: {
        marginRight: 10,
        width: 30,
        height: 30,
    },

    locationText: {
        color: "#3d3d3d",
    },  

    metadata: {
        marginTop: 20,
    },
    metadataText: {
        color: '#2557a7'
    },  

    jobDetails: {
        margin: 20
    },

    jobDescription: {
        backgroundColor: '#3d3d3d',
        paddingTop: 20,
        paddingHorizontal: 20,
    },

    bottomSheet: {
        backgroundColor: '#3d3d3d',
        paddingHorizontal: 20,
        paddingVertical: 20,
        borderTopLeftRadius: 20,
        borderTopRightRadius: 20,
        borderBottomWidth: 1,
        borderColor: 'grey'
    },

    descriptionText: {
        fontSize: 15,
        fontWeight: 'bold',
        color: '#fff'
    },

    jobDescriptionText: {
        color: "#fff"
    },


    tagContainer: {
        marginTop: 20,
        flexDirection: 'row',
        flexWrap: 'wrap'
    }
});