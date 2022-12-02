package builder
import . "builder/interface"



func getBuilder(builderType string) IBuilder {
    if builderType == "normal" {
        return newNormalBuilder()
    }

    if builderType == "igloo" {
        return newIglooBuilder()
    }
    return nil
}
