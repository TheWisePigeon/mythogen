use actix_web::{get, App, HttpResponse, HttpServer, Responder};

#[get("/")]
async fn mythogen() -> impl Responder {
    HttpResponse::Ok()
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    let command = std::env::args().nth(1);
    match command {
        Some(command) => match command.as_str() {
            "serve" => {
                if let Ok(port) = std::env::args()
                    .nth(2)
                    .unwrap_or(String::from("8080"))
                    .parse::<u16>()
                {
                    if port <= 1024 {
                        println!("You entered an invalid port! Make sure the port you are using is above 1024");
                        Ok(())
                    } else {
                        println!("Mythogen server launched on port {}", port);
                        HttpServer::new(|| App::new().service(mythogen))
                            .bind(("127.0.0.1", port))?
                            .run()
                            .await
                    }
                } else {
                    println!("You entered an invalid http port");
                    Ok(())
                }
            }
            "help" => Ok(()),
            _ => Ok(()),
        },
        None => Ok(()),
    }
}
