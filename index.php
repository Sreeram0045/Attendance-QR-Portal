<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>AM-QR</title>
    <link rel="stylesheet" href="./style/style.css">
</head>

<body class="scroll-smooth font-Poppins *:outline-none  focus-within:outline-none ">
    <div class="lg:w-screen lg:h-screen lg:flex items-center justify-center">
        <form action="" method="post" class="lg:flex lg:flex-row lg:gap-5 bg-[#FBF9E4] lg:p-20 lg:rounded-3xl">
            <div class="w-96 h-96">
                <img src="./assets/logo/logo.png" alt="logo" class="inline-block max-w-full max-h-full lg:rounded-full">
            </div>
            <main class="text-lg border">
                <div class="border">
                    <img src="./assets/icons/person_icon.svg" alt="person" class="opacity-50 inline-block w-8 h-10 mr-2">
                    <input type="text" name="f_name" id="full_name" placeholder="Enter Your Full Name" class=" align-middle bg-transparent">
                </div>
                <div class="border">
                    <input type="text" name="u_age" id="u_age" placeholder="Age" class="bg-transparent">
                </div>
                <div class="border">
                    <label>DOB:</label>
                    <input type="date" name="u_dob" id="u_dob" class="bg-transparent">
                </div>
                <div class="border">
                    <img src="./assets/icons/mail_icon.svg" alt="mail" class="opacity-50 inline-block w-8 h-9 mr-2">
                    <input type="email" name="u_mail" id="u_mail" placeholder="E-mail Id" class="align-middle bg-transparent">
                </div>
                <div class="border">
                    <img src="./assets/icons/phone.svg" alt="phone" class="opacity-50 inline-block w-8 h-9 mr-2">
                    <input type="text" name="u_num" id="u_num" placeholder="Phone Number" class="align-middle bg-transparent">
                </div>
                <div class="border">
                    <img src="./assets/icons/institution.png" alt="institution" class=" opacity-50 inline-block mr-2 h-8 w-7">
                    <input type="text" name="u_org" id="u_org" placeholder="Organisation" class="align-middle bg-transparent">
                </div>
                <div class="border">
                    <img src="./assets/icons/username.svg" alt="person" class="opacity-50 inline-block w-8 h-10 mr-2">
                    <input type="text" name="u_username" id="u_username" placeholder="Username" class="align-middle bg-transparent">
                </div>
                <div class="border">
                    <img src="./assets/icons/password.svg" alt="person" class="opacity-50 inline-block w-8 h-10 mr-2">
                    <input type="password" name="u_password" id="u_password" placeholder="Password" class="align-middle bg-transparent">
                </div>
                <div class="border">
                    <label>Upload your Photo:
                        <img src="./assets/icons/image_icon.svg" alt="image" class="opacity-50 inline-block w-8 h-9 mr-2">
                    </label>
                    <input type="file" name="u_img" id="u_img" class="bg-transparent">
                </div>
            </main>
        </form>
    </div>
</body>

</html>