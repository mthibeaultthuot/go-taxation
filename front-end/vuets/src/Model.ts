export interface Tax {
    Username: String,
    IsPstValid: boolean
    IsQstValid: boolean
    PstNumber: String
    QstNumber: String
    Enterprise: String
    Date: String
}


export interface TaxRequest {
    PstNumber: String
    QstNumber: String
}