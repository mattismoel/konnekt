import { Select, SelectContent, SelectGroup, SelectItem, SelectLabel, SelectTrigger, SelectValue } from "./select"

type Props = {
  defaultValue: string;
}

export const CountryPicker = ({ defaultValue }: Props) => {
  return (
    <Select name="country" defaultValue={defaultValue}>
      <SelectTrigger className="w-[280px]">
        <SelectValue placeholder="Vælg land..." />
      </SelectTrigger>
      <SelectContent>
        <SelectGroup  >
          <SelectLabel>Africa</SelectLabel>
          <SelectItem value="DZ" >Algeria</SelectItem>
          <SelectItem value="AO" >Angola</SelectItem>
          <SelectItem value="BJ" >Benin</SelectItem>
          <SelectItem value="BW" >Botswana</SelectItem>
          <SelectItem value="BF" >Burkina Faso</SelectItem>
          <SelectItem value="BI" >Burundi</SelectItem>
          <SelectItem value="CM" >Cameroon</SelectItem>
          <SelectItem value="CV" >Cape Verde</SelectItem>
          <SelectItem value="CF" >Central African Republic</SelectItem>
          <SelectItem value="TD" >Chad</SelectItem>
          <SelectItem value="KM" >Comoros</SelectItem>
          <SelectItem value="CG" >Congo - Brazzaville</SelectItem>
          <SelectItem value="CD" >Congo - Kinshasa</SelectItem>
          <SelectItem value="CI" >Côte d’Ivoire</SelectItem>
          <SelectItem value="DJ" >Djibouti</SelectItem>
          <SelectItem value="EG" >Egypt</SelectItem>
          <SelectItem value="GQ" >Equatorial Guinea</SelectItem>
          <SelectItem value="ER" >Eritrea</SelectItem>
          <SelectItem value="ET" >Ethiopia</SelectItem>
          <SelectItem value="GA" >Gabon</SelectItem>
          <SelectItem value="GM" >Gambia</SelectItem>
          <SelectItem value="GH" >Ghana</SelectItem>
          <SelectItem value="GN" >Guinea</SelectItem>
          <SelectItem value="GW" >Guinea-Bissau</SelectItem>
          <SelectItem value="KE" >Kenya</SelectItem>
          <SelectItem value="LS" >Lesotho</SelectItem>
          <SelectItem value="LR" >Liberia</SelectItem>
          <SelectItem value="LY" >Libya</SelectItem>
          <SelectItem value="MG" >Madagascar</SelectItem>
          <SelectItem value="MW" >Malawi</SelectItem>
          <SelectItem value="ML" >Mali</SelectItem>
          <SelectItem value="MR" >Mauritania</SelectItem>
          <SelectItem value="MU" >Mauritius</SelectItem>
          <SelectItem value="YT" >Mayotte</SelectItem>
          <SelectItem value="MA" >Morocco</SelectItem>
          <SelectItem value="MZ" >Mozambique</SelectItem>
          <SelectItem value="NA" >Namibia</SelectItem>
          <SelectItem value="NE" >Niger</SelectItem>
          <SelectItem value="NG" >Nigeria</SelectItem>
          <SelectItem value="RW" >Rwanda</SelectItem>
          <SelectItem value="RE" >Réunion</SelectItem>
          <SelectItem value="SH" >Saint Helena</SelectItem>
          <SelectItem value="SN" >Senegal</SelectItem>
          <SelectItem value="SC" >Seychelles</SelectItem>
          <SelectItem value="SL" >Sierra Leone</SelectItem>
          <SelectItem value="SO" >Somalia</SelectItem>
          <SelectItem value="ZA" >South Africa</SelectItem>
          <SelectItem value="SD" >Sudan</SelectItem>
          <SelectItem value="SZ" >Swaziland</SelectItem>
          <SelectItem value="ST" >São Tomé and Príncipe</SelectItem>
          <SelectItem value="TZ" >Tanzania</SelectItem>
          <SelectItem value="TG" >Togo</SelectItem>
          <SelectItem value="TN" >Tunisia</SelectItem>
          <SelectItem value="UG" >Uganda</SelectItem>
          <SelectItem value="EH" >Western Sahara</SelectItem>
          <SelectItem value="ZM" >Zambia</SelectItem>
          <SelectItem value="ZW" >Zimbabwe</SelectItem>
        </SelectGroup>
        <SelectGroup >
          <SelectLabel>Americas</SelectLabel>
          <SelectItem value="AI" >Anguilla</SelectItem>
          <SelectItem value="AG" >Antigua and Barbuda</SelectItem>
          <SelectItem value="AR" >Argentina</SelectItem>
          <SelectItem value="AW" >Aruba</SelectItem>
          <SelectItem value="BS" >Bahamas</SelectItem>
          <SelectItem value="BB" >Barbados</SelectItem>
          <SelectItem value="BZ" >Belize</SelectItem>
          <SelectItem value="BM" >Bermuda</SelectItem>
          <SelectItem value="BO" >Bolivia</SelectItem>
          <SelectItem value="BR" >Brazil</SelectItem>
          <SelectItem value="VG" >British Virgin Islands</SelectItem>
          <SelectItem value="CA" >Canada</SelectItem>
          <SelectItem value="KY" >Cayman Islands</SelectItem>
          <SelectItem value="CL" >Chile</SelectItem>
          <SelectItem value="CO" >Colombia</SelectItem>
          <SelectItem value="CR" >Costa Rica</SelectItem>
          <SelectItem value="CU" >Cuba</SelectItem>
          <SelectItem value="DM" >Dominica</SelectItem>
          <SelectItem value="DO" >Dominican Republic</SelectItem>
          <SelectItem value="EC" >Ecuador</SelectItem>
          <SelectItem value="SV" >El Salvador</SelectItem>
          <SelectItem value="FK" >Falkland Islands</SelectItem>
          <SelectItem value="GF" >French Guiana</SelectItem>
          <SelectItem value="GL" >Greenland</SelectItem>
          <SelectItem value="GD" >Grenada</SelectItem>
          <SelectItem value="GP" >Guadeloupe</SelectItem>
          <SelectItem value="GT" >Guatemala</SelectItem>
          <SelectItem value="GY" >Guyana</SelectItem>
          <SelectItem value="HT" >Haiti</SelectItem>
          <SelectItem value="HN" >Honduras</SelectItem>
          <SelectItem value="JM" >Jamaica</SelectItem>
          <SelectItem value="MQ" >Martinique</SelectItem>
          <SelectItem value="MX" >Mexico</SelectItem>
          <SelectItem value="MS" >Montserrat</SelectItem>
          <SelectItem value="AN" >Netherlands Antilles</SelectItem>
          <SelectItem value="NI" >Nicaragua</SelectItem>
          <SelectItem value="PA" >Panama</SelectItem>
          <SelectItem value="PY" >Paraguay</SelectItem>
          <SelectItem value="PE" >Peru</SelectItem>
          <SelectItem value="PR" >Puerto Rico</SelectItem>
          <SelectItem value="BL" >Saint Barthélemy</SelectItem>
          <SelectItem value="KN" >Saint Kitts and Nevis</SelectItem>
          <SelectItem value="LC" >Saint Lucia</SelectItem>
          <SelectItem value="MF" >Saint Martin</SelectItem>
          <SelectItem value="PM" >Saint Pierre and Miquelon</SelectItem>
          <SelectItem value="VC" >Saint Vincent and the Grenadines</SelectItem>
          <SelectItem value="SR" >Suriname</SelectItem>
          <SelectItem value="TT" >Trinidad and Tobago</SelectItem>
          <SelectItem value="TC" >Turks and Caicos Islands</SelectItem>
          <SelectItem value="VI" >U.S. Virgin Islands</SelectItem>
          <SelectItem value="US" >United States</SelectItem>
          <SelectItem value="UY" >Uruguay</SelectItem>
          <SelectItem value="VE" >Venezuela</SelectItem>
        </SelectGroup>
        <SelectGroup  >
          <SelectLabel>Asia</SelectLabel>
          <SelectItem value="AF" >Afghanistan</SelectItem>
          <SelectItem value="AM" >Armenia</SelectItem>
          <SelectItem value="AZ" >Azerbaijan</SelectItem>
          <SelectItem value="BH" >Bahrain</SelectItem>
          <SelectItem value="BD" >Bangladesh</SelectItem>
          <SelectItem value="BT" >Bhutan</SelectItem>
          <SelectItem value="BN" >Brunei</SelectItem>
          <SelectItem value="KH" >Cambodia</SelectItem>
          <SelectItem value="CN" >China</SelectItem>
          <SelectItem value="GE" >Georgia</SelectItem>
          <SelectItem value="HK" >Hong Kong SAR China</SelectItem>
          <SelectItem value="IN" >India</SelectItem>
          <SelectItem value="ID" >Indonesia</SelectItem>
          <SelectItem value="IR" >Iran</SelectItem>
          <SelectItem value="IQ" >Iraq</SelectItem>
          <SelectItem value="IL" >Israel</SelectItem>
          <SelectItem value="JP" >Japan</SelectItem>
          <SelectItem value="JO" >Jordan</SelectItem>
          <SelectItem value="KZ" >Kazakhstan</SelectItem>
          <SelectItem value="KW" >Kuwait</SelectItem>
          <SelectItem value="KG" >Kyrgyzstan</SelectItem>
          <SelectItem value="LA" >Laos</SelectItem>
          <SelectItem value="LB" >Lebanon</SelectItem>
          <SelectItem value="MO" >Macau SAR China</SelectItem>
          <SelectItem value="MY" >Malaysia</SelectItem>
          <SelectItem value="MV" >Maldives</SelectItem>
          <SelectItem value="MN" >Mongolia</SelectItem>
          <SelectItem value="MM" >Myanmar [Burma]</SelectItem>
          <SelectItem value="NP" >Nepal</SelectItem>
          <SelectItem value="NT" >Neutral Zone</SelectItem>
          <SelectItem value="KP" >North Korea</SelectItem>
          <SelectItem value="OM" >Oman</SelectItem>
          <SelectItem value="PK" >Pakistan</SelectItem>
          <SelectItem value="PS" >Palestinian Territories</SelectItem>
          <SelectItem value="YD" >People's Democratic Republic of Yemen</SelectItem>
          <SelectItem value="PH" >Philippines</SelectItem>
          <SelectItem value="QA" >Qatar</SelectItem>
          <SelectItem value="SA" >Saudi Arabia</SelectItem>
          <SelectItem value="SG" >Singapore</SelectItem>
          <SelectItem value="KR" >South Korea</SelectItem>
          <SelectItem value="LK" >Sri Lanka</SelectItem>
          <SelectItem value="SY" >Syria</SelectItem>
          <SelectItem value="TW" >Taiwan</SelectItem>
          <SelectItem value="TJ" >Tajikistan</SelectItem>
          <SelectItem value="TH" >Thailand</SelectItem>
          <SelectItem value="TL" >Timor-Leste</SelectItem>
          <SelectItem value="TR" >Turkey</SelectItem>
          <SelectItem value="TM" >Turkmenistan</SelectItem>
          <SelectItem value="AE" >United Arab Emirates</SelectItem>
          <SelectItem value="UZ" >Uzbekistan</SelectItem>
          <SelectItem value="VN" >Vietnam</SelectItem>
          <SelectItem value="YE" >Yemen</SelectItem>
        </SelectGroup>
        <SelectGroup  >
          <SelectLabel>Europe</SelectLabel>
          <SelectItem value="AL" >Albania</SelectItem>
          <SelectItem value="AD" >Andorra</SelectItem>
          <SelectItem value="AT" >Austria</SelectItem>
          <SelectItem value="BY" >Belarus</SelectItem>
          <SelectItem value="BE" >Belgium</SelectItem>
          <SelectItem value="BA" >Bosnia and Herzegovina</SelectItem>
          <SelectItem value="BG" >Bulgaria</SelectItem>
          <SelectItem value="HR" >Croatia</SelectItem>
          <SelectItem value="CY" >Cyprus</SelectItem>
          <SelectItem value="CZ" >Czech Republic</SelectItem>
          <SelectItem value="DK" >Denmark</SelectItem>
          <SelectItem value="DD" >East Germany</SelectItem>
          <SelectItem value="EE" >Estonia</SelectItem>
          <SelectItem value="FO" >Faroe Islands</SelectItem>
          <SelectItem value="FI" >Finland</SelectItem>
          <SelectItem value="FR" >France</SelectItem>
          <SelectItem value="DE" >Germany</SelectItem>
          <SelectItem value="GI" >Gibraltar</SelectItem>
          <SelectItem value="GR" >Greece</SelectItem>
          <SelectItem value="GG" >Guernsey</SelectItem>
          <SelectItem value="HU" >Hungary</SelectItem>
          <SelectItem value="IS" >Iceland</SelectItem>
          <SelectItem value="IE" >Ireland</SelectItem>
          <SelectItem value="IM" >Isle of Man</SelectItem>
          <SelectItem value="IT" >Italy</SelectItem>
          <SelectItem value="JE" >Jersey</SelectItem>
          <SelectItem value="LV" >Latvia</SelectItem>
          <SelectItem value="LI" >Liechtenstein</SelectItem>
          <SelectItem value="LT" >Lithuania</SelectItem>
          <SelectItem value="LU" >Luxembourg</SelectItem>
          <SelectItem value="MK" >Macedonia</SelectItem>
          <SelectItem value="MT" >Malta</SelectItem>
          <SelectItem value="FX" >Metropolitan France</SelectItem>
          <SelectItem value="MD" >Moldova</SelectItem>
          <SelectItem value="MC" >Monaco</SelectItem>
          <SelectItem value="ME" >Montenegro</SelectItem>
          <SelectItem value="NL" >Netherlands</SelectItem>
          <SelectItem value="NO" >Norway</SelectItem>
          <SelectItem value="PL" >Poland</SelectItem>
          <SelectItem value="PT" >Portugal</SelectItem>
          <SelectItem value="RO" >Romania</SelectItem>
          <SelectItem value="RU" >Russia</SelectItem>
          <SelectItem value="SM" >San Marino</SelectItem>
          <SelectItem value="RS" >Serbia</SelectItem>
          <SelectItem value="CS" >Serbia and Montenegro</SelectItem>
          <SelectItem value="SK" >Slovakia</SelectItem>
          <SelectItem value="SI" >Slovenia</SelectItem>
          <SelectItem value="ES" >Spain</SelectItem>
          <SelectItem value="SJ" >Svalbard and Jan Mayen</SelectItem>
          <SelectItem value="SE" >Sweden</SelectItem>
          <SelectItem value="CH" >Switzerland</SelectItem>
          <SelectItem value="UA" >Ukraine</SelectItem>
          <SelectItem value="SU" >Union of Soviet Socialist Republics</SelectItem>
          <SelectItem value="GB" >United Kingdom</SelectItem>
          <SelectItem value="VA" >Vatican City</SelectItem>
          <SelectItem value="AX" >Åland Islands</SelectItem>
        </SelectGroup>
        <SelectGroup  >
          <SelectLabel>Oceania</SelectLabel>
          <SelectItem value="AS" >American Samoa</SelectItem>
          <SelectItem value="AQ" >Antarctica</SelectItem>
          <SelectItem value="AU" >Australia</SelectItem>
          <SelectItem value="BV" >Bouvet Island</SelectItem>
          <SelectItem value="IO" >British Indian Ocean Territory</SelectItem>
          <SelectItem value="CX" >Christmas Island</SelectItem>
          <SelectItem value="CC" >Cocos [Keeling] Islands</SelectItem>
          <SelectItem value="CK" >Cook Islands</SelectItem>
          <SelectItem value="FJ" >Fiji</SelectItem>
          <SelectItem value="PF" >French Polynesia</SelectItem>
          <SelectItem value="TF" >French Southern Territories</SelectItem>
          <SelectItem value="GU" >Guam</SelectItem>
          <SelectItem value="HM" >Heard Island and McDonald Islands</SelectItem>
          <SelectItem value="KI" >Kiribati</SelectItem>
          <SelectItem value="MH" >Marshall Islands</SelectItem>
          <SelectItem value="FM" >Micronesia</SelectItem>
          <SelectItem value="NR" >Nauru</SelectItem>
          <SelectItem value="NC" >New Caledonia</SelectItem>
          <SelectItem value="NZ" >New Zealand</SelectItem>
          <SelectItem value="NU" >Niue</SelectItem>
          <SelectItem value="NF" >Norfolk Island</SelectItem>
          <SelectItem value="MP" >Northern Mariana Islands</SelectItem>
          <SelectItem value="PW" >Palau</SelectItem>
          <SelectItem value="PG" >Papua New Guinea</SelectItem>
          <SelectItem value="PN" >Pitcairn Islands</SelectItem>
          <SelectItem value="WS" >Samoa</SelectItem>
          <SelectItem value="SB" >Solomon Islands</SelectItem>
          <SelectItem value="GS" >South Georgia and the South Sandwich Islands</SelectItem>
          <SelectItem value="TK" >Tokelau</SelectItem>
          <SelectItem value="TO" >Tonga</SelectItem>
          <SelectItem value="TV" >Tuvalu</SelectItem>
          <SelectItem value="UM" >U.S. Minor Outlying Islands</SelectItem>
          <SelectItem value="VU" >Vanuatu</SelectItem>
          <SelectItem value="WF" >Wallis and Futuna</SelectItem>
        </SelectGroup>
      </SelectContent>
    </Select >
  )
}
